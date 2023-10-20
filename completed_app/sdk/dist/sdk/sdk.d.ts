import * as utils from "../internal/utils";
import { Band } from "./band";
import { DrinkOperations } from "./drinkoperations";
import { MusicOperations } from "./musicoperations";
import { AxiosInstance } from "axios";
/**
 * Contains the list of servers available to the SDK
 */
export declare const ServerList: readonly ["http://{machine}:51000"];
/**
 * The available configuration options for the SDK
 */
export type SDKProps = {
    /**
     * Allows overriding the default axios client used by the SDK
     */
    defaultClient?: AxiosInstance;
    /**
     * Allows overriding the default server used by the SDK
     */
    serverIdx?: number;
    /**
     * Allows setting the machine variable for url substitution
     */
    machine?: string;
    /**
     * Allows overriding the default server URL used by the SDK
     */
    serverURL?: string;
    /**
     * Allows overriding the default retry config used by the SDK
     */
    retryConfig?: utils.RetryConfig;
};
export declare class SDKConfiguration {
    defaultClient: AxiosInstance;
    serverURL: string;
    serverDefaults: any;
    language: string;
    openapiDocVersion: string;
    sdkVersion: string;
    genVersion: string;
    userAgent: string;
    retryConfig?: utils.RetryConfig;
    constructor(init?: Partial<SDKConfiguration>);
}
/**
 * The Speakeasy Club: A club that serves drinks and plays jazz. A Goa and Speakeasy example.
 */
export declare class SDK {
    drinkOperations: DrinkOperations;
    musicOperations: MusicOperations;
    /**
     * A band that plays jazz.
     */
    band: Band;
    private sdkConfiguration;
    constructor(props?: SDKProps);
}
