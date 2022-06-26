import React, { useState, useEffect } from 'react';
import './App.css';
import { User } from './types';
import { NoAuth } from './pages/NoAuth';
import { Auth } from './pages/Auth';
import axios from 'axios';


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
        } else {
          return Promise.reject(response);
        }
      }).catch(e => {
        console.log(e);
      })
  }, [])
  return (
    <div>
      Ledger App
      {loggedIn ? <Auth user={userInfo} /> : <NoAuth />}
    </div>
  );
}

export default App;
