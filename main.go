package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	appName = "untilitworks"
	version = "devel"
)

func main() {

	retryType := flag.String("retry", "constantly", "retry type, one of 'c'/'constantly' or 'e'/'exponentially'")
	sleep := flag.Duration("sleep", time.Second, "how long to sleep between retries (base duration for exponential)")
	max := flag.Duration("max", 0, "it set, maximum duration for which to wait until it works")
	factor := flag.Float64("exp.factor", 2, "backoff factor for exponential retries")
	cap := flag.Duration("exp.cap", 30*time.Second, "max time to backoff for")
	quiet := flag.Bool("q", false, "whether to suppress the command's output")
	pversion := flag.Bool("v", false, "print the version and exit")
	flag.Parse()

	if *pversion {
		fmt.Println(appName + "-" + version)
		os.Exit(0)
	}

	log.SetFlags(0)
	log.SetPrefix(appName + ": ")
	if *quiet {
		log.SetOutput(ioutil.Discard)
	}

	var retryFunc func()

	switch *retryType {
	case "c", "constantly":

		retryFunc = func() {
			if *sleep != 0 {
				log.Printf("it failed! retrying in %v", *sleep)
				time.Sleep(*sleep)
			} else {
				log.Printf("it failed!")
			}
		}

	case "e", "exponentially":

		if *sleep <= 0 {
			log.Fatal("can't have exponential retries with 0 second base sleep")
		}
		if *factor <= 0 {
			log.Fatal("can't have exponential backoff factor of 0")
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		baseSeconds := sleep.Seconds()
		capSecond := cap.Seconds()
		attempt := 0.0

		retryFunc = func() {
			sleepSeconds := r.Float64() * math.Min(capSecond, baseSeconds*math.Pow(*factor, attempt))
			sleep := time.Duration(sleepSeconds * float64(time.Second))
			log.Printf("it failed! retrying in %v", sleep)
			time.Sleep(sleep)
			attempt += 1.0
		}

	default:
		flag.PrintDefaults()
		log.Fatalf("not a valid retry type: %q", *retryType)
	}

	cmds := flag.Args()

	if len(cmds) == 0 {
		log.Fatal("need a command to run")
	}

	start := time.Now()
	for {
		cmd := exec.Command(cmds[0], cmds[1:]...)
		if !*quiet {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
		if err := cmd.Run(); err == nil {
			break
		}
		if *max != 0 && time.Since(start) >= *max {
			log.Fatalf("it never worked!")
		}
		retryFunc()
	}

	log.Printf("it worked!")
}
