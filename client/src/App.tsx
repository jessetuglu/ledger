import React, { useState, useEffect } from 'react';
import './App.css';
import { User } from './types';
import { NoAuth } from './pages/NoAuth';
import { Auth } from './pages/Auth';
import axios from 'axios';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Ledger } from './components/Ledger';
import { Transaction } from './components/Transaction';
import { Main } from './components/Main';


function App() {
  const [userInfo, setUserInfo] = useState<User | undefined>(undefined);
  const [loggedIn, setLoggedIn] = useState<boolean>(false);


  useEffect(() => {
    axios.get("http://localhost:8080/api/auth/whoami", { withCredentials: true })
      .then(response => {
        if (response.status == 200) {
          console.log(response.data);
          const data = response.data;
          setLoggedIn(true);
          setUserInfo({ email: data.Email, first_name: data.FirstName, last_name: data.LastName, id: data.ID });
        };
      }).catch(e => {
        console.log(e);
        setLoggedIn(false);
        setUserInfo(undefined);
      })
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Main/>}/>
        {loggedIn ?
          <>
          <Route path="/ledgers/:id" element={<Ledger/>}/>
          <Route path="/transactions/:id" element={<Transaction/>}/>
          </>
          :
          null
        }
      </Routes>
    </BrowserRouter>
  );
}

export default App;


// flow
// 1. user login
// 2. if auth expose /home
// 3. else remain on homepage with login button
// 4. enforce no nav to /home with global state
// 5. once on home, make a request to DB to retrieve list of ledgers render these in list group
// 6. once a user clicks on a ledger, make another request to fetch ledger transactions to display on screen (Ledger component)