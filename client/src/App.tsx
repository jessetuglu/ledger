import React, {useState, useEffect} from 'react';
import './App.css';
import {User} from './types/User';
import { NoAuth } from './views/NoAuth';
import { Auth } from './views/Auth';


function App() {
  const [userInfo, setUserInfo] = useState<User|undefined>(undefined);
  const [loggedIn, setLoggedIn] = useState<boolean>(false);
  
  
  useEffect(() => {
    fetch("http://localhost:8080/api/auth/whoami") //TODO: make this ENV
    .then(response => {
      if (response.status == 200){
        response.json();
      }else{
        return Promise.reject(response);
      }
    }).then(data => {
      console.log(data);
      
      setLoggedIn(true);
      setUserInfo(data);
    })
    .catch(e => {
      console.log("User is not authenticated.");
      console.log(e)
    })
  },[])
  return (
    <div>
      Ledger App
      {loggedIn ? <Auth/> : <NoAuth/>}
    </div>
  );
}

export default App;
