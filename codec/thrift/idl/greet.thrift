namespace go proxyless

struct HelloRequest {
    1: required string Message,
}

struct HelloResponse {
    1: required string Message,
}

service GreetService {
    HelloResponse SayHello(1: HelloRequest request);
    HelloResponse SayHi(1: HelloRequest request);
}
