import { SpeakeasyBase } from "../../../internal/utils";
export declare class TeaRequestBody extends SpeakeasyBase {
    /**
     * Whether to have milk.
     */
    includeMilk?: boolean;
    /**
     * Whether to have green tea instead of normal.
     */
    isGreen?: boolean;
    /**
     * Number of spoons of sugar.
     */
    numberSugars?: number;
}
