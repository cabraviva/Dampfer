import type { IconDefinition } from "@fortawesome/free-solid-svg-icons"

type ColorType = "gray" | "red" | "yellow" | "green" | "indigo" | "purple" | "pink" | "blue" | "light" | "dark" | "default" | "dropdown" | "navbar" | "navbarUl" | "form" | "primary" | "orange" | "none"

export interface GeneralAlert {
    bold: string
    msg: string,
    color: ColorType
}

export interface AlertType {
    /**
     * Icon to show. Must be a font awesome icon
     */
    icon: IconDefinition

    color: ColorType

    content: string

    /**
     * Whether the alert should have a view button. You can also define a onclick handler and a custom text
     * @default false (disabled)
     */
    viewButton?: false | {
        onclick: Function,
        text?: string
    }

    /**
     * Whether the alert can be closed, or if it should respawn
     * @default true
     */
    canBeIgnored?: boolean
}