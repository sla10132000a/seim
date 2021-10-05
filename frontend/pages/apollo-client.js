import { ApolloClient, InMemoryCache } from "@apollo/client";

const client = new ApolloClient({
    // uri: "https://countries.trevorblades.com",
    uri: "http://localhost:8080/query",
    cache: new InMemoryCache(),
    credentials: 'same-origin',
    ssrMode: typeof window === 'undefined',
});

export default client;