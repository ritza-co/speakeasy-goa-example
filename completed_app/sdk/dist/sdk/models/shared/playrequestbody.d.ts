import { SpeakeasyBase } from "../../../internal/utils";
/**
 * Style of music to play
 */
export declare enum PlayRequestBodyStyle {
    Bebop = "Bebop",
    Swing = "Swing"
}
export declare class PlayRequestBody extends SpeakeasyBase {
    /**
     * Style of music to play
     */
    style: PlayRequestBodyStyle;
}
