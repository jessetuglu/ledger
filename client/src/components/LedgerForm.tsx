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

    const deleteLedger = (event: any) => {
        event.preventDefault();
        axios.delete("http://localhost:8080/api/ledgers/"+event.target.name)
        .then((resp)=>{
            console.log(resp);
        })
        .catch((e)=>{
            console.log(e);
        });
    }

    useEffect(() => {
        axios.get("http://localhost:8080/api/users/"+user?.id+"/ledgers", {withCredentials: true})
        .then((resp) => {
            console.log(resp.data);
            setLedgers(resp.data);
        })
    }, [createForm]);
    return (
        <div>
            <h3>Ledgers</h3>
            <ul className="list-group">
                {ledgers.map((ledger:any, key)=> {
                    return (
                        <li className="list-group-item">
                            <p>{ledger.Title}</p>
                            <button name={ledger.ID} onClick={deleteLedger}>Delete</button>
                        </li>
                    );
                })}
            </ul>
            <input type={"text"} onChange={setFormHelper}></input>
            <button onClick={createForm}>Submit</button>
        </div>
    )
}