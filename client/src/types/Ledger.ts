import { User } from "./User"

export type Ledger = {
    id: string
    title: string
    members: [User]
}