export type Transaction = {
    id: string;
    amount: number;
    debitor: string;
    creditor: string;
    note: string;
    date: string;
};

export type Ledger = {
    id: string;
    title: string;
    members: [User];
};

export type User = {
    email: string;
    first_name: string;
    last_name: string;
    id: string;
};