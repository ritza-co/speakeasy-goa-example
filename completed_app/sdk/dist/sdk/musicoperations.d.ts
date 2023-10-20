import * as operations from "./models/operations";
import * as shared from "./models/shared";
import { SDKConfiguration } from "./sdk";
import { AxiosRequestConfig } from "axios";
export declare class MusicOperations {
    private sdkConfiguration;
    constructor(sdkConfig: SDKConfiguration);
    /**
     * play band
     *
     * @remarks
     * Choose your jazz style.
     */
    bandNumberPlay(req: shared.PlayRequestBody, config?: AxiosRequestConfig): Promise<operations.BandNumberPlayResponse>;
}
