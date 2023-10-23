/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as utils from "../internal/utils";
import { Band } from "./band";
import { DrinkOperations } from "./drinkoperations";
import { MusicOperations } from "./musicoperations";
import axios from "axios";
import { AxiosInstance } from "axios";

/**
 * Contains the list of servers available to the SDK
 */
export const ServerList = [
    /**
     * club server hosts the band and order services.
     */
    "http://{machine}:51000",
] as const;

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

export class SDKConfiguration {
    defaultClient: AxiosInstance;
    serverURL: string;
    serverDefaults: any;
    language = "typescript";
    openapiDocVersion = "1.0.0";
    sdkVersion = "0.0.1";
    genVersion = "2.169.0";
    userAgent = "speakeasy-sdk/typescript 0.0.1 2.169.0 1.0.0 openapi";
    retryConfig?: utils.RetryConfig;
    public constructor(init?: Partial<SDKConfiguration>) {
        Object.assign(this, init);
    }
}

/**
 * The Speakeasy Club: A club that serves drinks and plays jazz. A Goa and Speakeasy example.
 */
export class SDK {
    public drinkOperations: DrinkOperations;
    public musicOperations: MusicOperations;
    /**
     * A band that plays jazz.
     */
    public band: Band;

    private sdkConfiguration: SDKConfiguration;

    constructor(props?: SDKProps) {
        let serverURL = props?.serverURL;
        let defaults: any = {};

        const serverDefaults = [
            {
                machine: props?.machine?.toString() ?? "localhost",
            },
        ];
        const serverIdx = props?.serverIdx ?? 0;

        if (!serverURL) {
            serverURL = ServerList[serverIdx];
            defaults = serverDefaults[serverIdx];
        }

        const defaultClient = props?.defaultClient ?? axios.create({ baseURL: serverURL });
        this.sdkConfiguration = new SDKConfiguration({
            defaultClient: defaultClient,
            serverURL: serverURL,
            serverDefaults: defaults,
            retryConfig: props?.retryConfig,
        });

        this.drinkOperations = new DrinkOperations(this.sdkConfiguration);
        this.musicOperations = new MusicOperations(this.sdkConfiguration);
        this.band = new Band(this.sdkConfiguration);
    }
}
