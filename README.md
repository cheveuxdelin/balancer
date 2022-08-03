# BALANCER🐈‍⬛

![cat](https://media.giphy.com/media/myDXHnYYPxT4od4NEN/giphy-downsized.gif)

## What is this?

This is a Concurrent design pattern made in Go, reusable for many projects, which I called Resource Balancer!

Having some defined limited **resources,** (an example in real life could be technical support operators), and some continuous consumers for the resources called **units** (Following the analogy, these would be users requesting technical support over the phone), This design pattern assigns the consumers to available resources (distributing the calls over the available operators), also making them available again as soon as units stop requiring them, working in a multi-tasking nature.

## Sample run

![ezgif-4-66855cba23](https://user-images.githubusercontent.com/68461123/182526481-df59ee5e-ad60-4a26-81d2-07d4141144d7.gif)

Red represents resources being used and green represent available resources

## How is this helpful?

Lots of ways! I can think of a lot of cases where assigning limited resources to continuous consumers is a great thing: Web servers, task scheduling, uber drive assigning, the list can go on!

## How can this project grow?

I'm thinking on making some practical examples using the design pattern and showcasing them in this repo! Make sure to watch the repo to see where does this go🐝
