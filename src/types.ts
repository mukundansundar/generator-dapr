// Supported dapr modes
type DaprMode = "Kubernetes" | "Standalone";

// Supported languages
export type Language = "JavaScript" | "TypeScript" | "Go" | "Python" | "C#";

// Supported state stores
type StateStore = "Redis" | "CosmosDB" | "Cassandra" | "None";

// Supported protocols
type DaprProtocol = "HTTP" | "gRPC";

// Supported pubsub components
type PubSub = "Redis";

export interface Answers {
    name: string,
    mode: DaprMode,
    languages: Language[],
    store: string
}

export interface Microservice {
    /** The language of the microservice. */
    language: Language;
    /** Informs whether or not the microservice gets or sets state. */
    statePersistance?: boolean;
    /** Informs whether or not the microservice publishes or subscribes messages using pubsub. */
    pubsub?: boolean;
    /** Informs whether or not the microservice's manifest should include a LoadBalancer. */
    externalEndpoint?: boolean;
    /** Informs whether or not to add boilerplate actor code to microservice. */
    actors?: boolean;
    /** The protocol the microervice uses. */
    protocol?: DaprProtocol;
}

export interface App {
    name: string;
    /** Indicates whether the dapr app is running in Kubernetes or in Standalone mode. */
    mode: DaprMode;
    /** The microservices that compose the dapr application. */
    microservices: Microservice[];
    /** The state store to be used by the dapr application. */
    stateStore?: StateStore;
    /** The pubsub mechanism to be used by the dapr application. */
    pubsub?: PubSub;
}

export interface Prompt {
    type: string,
    name: string,
    message: string,
    choices?: string[]
}

export type Prompts = Prompt[];