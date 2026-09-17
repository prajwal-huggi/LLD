package main

import (
	"context"
	"fmt"
	"time"
	"github.com/go-redis/redis/v8"
)

type Tier struct{
	id string
	name string
	windowSize int
	requests int64
	algoName string
}

type Client struct{
	id string
	name string
	tier Tier
}

type User struct{
	id string
	name string
	client Client
}

type IStore interface{
	Incr(ctx context.Context, key string, ttl time.Duration)(int64, error)
}

type Redis struct{
	client *redis.Client
}
func(r *Redis) Incr(ctx context.Context, key string, ttl time.Duration)(int64, error){
	// fmt.Println("Getting the Redis")
	// get value from the redis key
	// increment the value in the redis
	// if there is no error return
	// else return the updated value
	// getVal:= 
	// increment the valuectx
	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println("Redis is down")
		return 0, fmt.Errorf("incr %s: %w", key, err)
	}

	if count == 1 {
		// first request in this window — start the clock
		if err := r.client.Expire(ctx, key, ttl).Err(); err != nil {
			return count, fmt.Errorf("expire %s: %w", key, err)
		}
	}

	return count, nil
}

type Decision struct{
	Allowed    bool
	RetryAfter time.Duration
}

type IRateLimitAlgo interface{
	ImplementAlgo(ctx context.Context, req string, tier Tier) (Decision, error)
}

type FixedWindow struct{
	store IStore
}
func(fw *FixedWindow) ImplementAlgo(ctx context.Context, req string, tier Tier)(Decision, error){
	fmt.Println("implement fixed window algo ", tier, " this is the tier")
	// fmt.Println("Either fail or pass")
	window:= time.Duration(tier.windowSize)* time.Second//seconds
	getValue, err:= fw.store.Incr(ctx, req, window) 
	if err!= nil{
		return Decision{
			Allowed: true,
		}, nil
	}
	
	totalReq:= tier.requests

	if getValue> totalReq{
		return Decision{
			Allowed: false,
			RetryAfter: window,
		}, nil
	}
	return Decision{
		Allowed: true,
	}, nil
}

// type SlidingWindow struct{
// 	store IStore
// }
// func(fw *SlidingWindow) implementAlgo(ctx context.Context, req string, tier Tier)(Decision, error){
// 	fmt.Println("sliding window algo ", tier, " this is the tier")
// 	fmt.Println("Either fail or pass")

// 	window:= ;
// 	getVal, err:= fw.store.Incr(ctx, req, window); if err!= nil{
// 		return Decision{Allowed: true}, nil
// 	}

// 	totalReq:= tier.requests

// 	return Decision{Allowed: true}, nil
// }


type IRateLimitFactory interface{
	createAlgo(store IStore) IRateLimitAlgo
}

type FixedWindowFactory struct{}
func(f* FixedWindowFactory) createAlgo(store IStore)IRateLimitAlgo{
	return &FixedWindow{
		store: store,
	}
}

// type SlidingWindowFactory struct{}
// func(f* SlidingWindowFactory) createAlgo(store IStore)IRateLimitAlgo{
// 	return &SlidingWindow{
// 		store: store,
// 	}
// }

type RateLimiterService struct{
	store     IStore
	factories map[string]IRateLimitFactory
	algos     map[string]IRateLimitAlgo // built once, read-only afterwards
}

func NewRateLimiterService(store IStore, factories map[string]IRateLimitFactory) *RateLimiterService{
	algos:= make(map[string]IRateLimitAlgo, len(factories))
	for name, f:= range factories{
		algos[name]= f.createAlgo(store)
	}
	return &RateLimiterService{
		store:     store,
		factories: factories,
		algos:     algos,
	}
}

func(r* RateLimiterService) GetAlgo(algoName string)(IRateLimitAlgo, error){
	algo, ok:= r.algos[algoName]
	if !ok{
		return nil, fmt.Errorf("unknown rate limit algo %q", algoName)
	}
	return algo, nil
}

func(r* RateLimiterService) Allow(ctx context.Context, user User)(Decision, error){
	tier:= user.client.tier

	algo, err:= r.GetAlgo(tier.algoName)
	if err!= nil{
		return Decision{}, err
	}

	key:= fmt.Sprintf("rl:%s:%s:%s", tier.algoName, tier.id, user.client.id)
	return algo.ImplementAlgo(ctx, key, tier)
}


func main(){
	store:= &Redis{
		client: redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
	}

	svc:= NewRateLimiterService(store, map[string]IRateLimitFactory{
		"fixed_window": &FixedWindowFactory{},
		// "sliding_window": &SlidingWindowFactory{},
	})

	user:= User{
		id: "u1", name: "prajwal",
		client: Client{
			id: "amz", name: "Amazon",
			tier: Tier{id: "t1", name: "free", windowSize: 60, requests: 2, algoName: "fixed_window"},
		},
	}

	decision, err:= svc.Allow(context.Background(), user)
	if err!= nil{
		fmt.Println("rate limit error:", err)
		return
	}
	fmt.Printf("allowed=%v retryAfter=%v\n", decision.Allowed, decision.RetryAfter)
	decision1, err:= svc.Allow(context.Background(), user)
	if err!= nil{
		fmt.Println("rate limit error:", err)
		return
	}
	fmt.Printf("allowed=%v retryAfter=%v\n", decision1.Allowed, decision1.RetryAfter)
	decision2, err:= svc.Allow(context.Background(), user)
	if err!= nil{
		fmt.Println("rate limit error:", err)
		return
	}
	fmt.Printf("allowed=%v retryAfter=%v\n", decision2.Allowed, decision2.RetryAfter)
	decision3, err:= svc.Allow(context.Background(), user)
	if err!= nil{
		fmt.Println("rate limit error:", err)
		return
	}
	fmt.Printf("allowed=%v retryAfter=%v\n", decision3.Allowed, decision3.RetryAfter)
}
