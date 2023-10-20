"use strict";
/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.SDK = exports.SDKConfiguration = exports.ServerList = void 0;
var band_1 = require("./band");
var drinkoperations_1 = require("./drinkoperations");
var musicoperations_1 = require("./musicoperations");
var axios_1 = __importDefault(require("axios"));
/**
 * Contains the list of servers available to the SDK
 */
exports.ServerList = [
    /**
     * club server hosts the band and order services.
     */
    "http://{machine}:51000",
];
var SDKConfiguration = /** @class */ (function () {
    function SDKConfiguration(init) {
        this.language = "typescript";
        this.openapiDocVersion = "1.0.0";
        this.sdkVersion = "0.0.1";
        this.genVersion = "2.155.1";
        this.userAgent = "speakeasy-sdk/typescript 0.0.1 2.155.1 1.0.0 openapi";
        Object.assign(this, init);
    }
    return SDKConfiguration;
}());
exports.SDKConfiguration = SDKConfiguration;
/**
 * The Speakeasy Club: A club that serves drinks and plays jazz. A Goa and Speakeasy example.
 */
var SDK = /** @class */ (function () {
    function SDK(props) {
        var _a, _b, _c, _d;
        var serverURL = props === null || props === void 0 ? void 0 : props.serverURL;
        var defaults = {};
        var serverDefaults = [
            {
                machine: (_b = (_a = props === null || props === void 0 ? void 0 : props.machine) === null || _a === void 0 ? void 0 : _a.toString()) !== null && _b !== void 0 ? _b : "localhost",
            },
        ];
        var serverIdx = (_c = props === null || props === void 0 ? void 0 : props.serverIdx) !== null && _c !== void 0 ? _c : 0;
        if (!serverURL) {
            serverURL = exports.ServerList[serverIdx];
            defaults = serverDefaults[serverIdx];
        }
        var defaultClient = (_d = props === null || props === void 0 ? void 0 : props.defaultClient) !== null && _d !== void 0 ? _d : axios_1.default.create({ baseURL: serverURL });
        this.sdkConfiguration = new SDKConfiguration({
            defaultClient: defaultClient,
            serverURL: serverURL,
            serverDefaults: defaults,
            retryConfig: props === null || props === void 0 ? void 0 : props.retryConfig,
        });
        this.drinkOperations = new drinkoperations_1.DrinkOperations(this.sdkConfiguration);
        this.musicOperations = new musicoperations_1.MusicOperations(this.sdkConfiguration);
        this.band = new band_1.Band(this.sdkConfiguration);
    }
    return SDK;
}());
exports.SDK = SDK;