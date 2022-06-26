import axios from "axios";
import React, { ChangeEvent, MouseEventHandler, useEffect, useState } from "react";
import { Ledger, User } from "../types";

interface UserProps {
    user: User|undefined,
}

export const LedgerForm: React.FC<UserProps> = ({ user }): JSX.Element => {
    const [ledgers, setLedgers] = useState([]);
    const [formTitle, setFormTitle] = useState("");

    const setFormHelper = (event: ChangeEvent<HTMLInputElement>) => {
        setFormTitle(event.target.value);
    };
    const createForm = (event: any) => {
        event.preventDefault();
        const body = {
            title: formTitle,
            members: [user?.id]
        };
        axios.post("http://localhost:8080/api/ledgers", body, {withCredentials: true})
        .then(resp => {
            console.log(resp);
        })
        .catch(e => {
            console.log(e);
        });
    };

    useEffect(() => {
        axios.get("http://localhost:8080/api/users/"+user?.id+"/ledgers", {withCredentials: true})
        .then((resp) => {
            console.log(resp.data);
            setLedgers(resp.data);
        })
    }, []);
    return (
        <div>
            <div>
                {ledgers.map((ledger:any, key)=> {
                    return (
                        <div key={key}>
                            <p>{ledger.ID}</p>
                            <p>{ledger.Title}</p>
                        </div>
                    );
                })}
            </div>
            <input type={"text"} onChange={setFormHelper}></input>
            <button onClick={createForm}>Submit</button>
        </div>
    )
}