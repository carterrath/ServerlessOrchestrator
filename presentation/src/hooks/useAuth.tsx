import { IUserData } from "../types/user-data";
import { useState, useEffect } from "react";

type UserType = 'Developer' | 'Consumer';

export function useAuth() {
  const [userDetails, setUserDetails] = useState<IUserData | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [clearStatesTimeout, setClearStatesTimeout] = useState<number | null>(null);

  async function login(username: string, password: string, userType: UserType) {
    try{
      const response = await fetch('http://localhost:8080/login/developer', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            username: username,
            password: password,
            userType: userType,
          }),
        });
        if (response.ok) {
          const responseData = await response.json();
          
          storeTokenInLocalStorage(responseData.token);
          console.log(responseData);
          fetchUserDetails();
          return "success";
          // navigate('/Home'); // Navigate on success
        } else {
          const errorData = await response.json();
          return errorData.error;
        }
    } catch (error) {
      return error;
    }
  }

  function storeTokenInLocalStorage(token: string) {
    localStorage.setItem('token', token); // Store token in local storage
  }

  async function fetchUserDetails() {
    const token = localStorage.getItem('token');
    if(!token) {
      // maybe redirect to login page
      console.log("No token found");
      return;
    }
    try{
      const response = await fetch('http://localhost:8080/getuserdetails', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}` // Add the token to the Authorization header
      },
      credentials: 'include'
      });

      if (!response.ok) {
        throw new Error('Failed to fetch user details');
      }
      console.log("response ok");
      // If response is OK, parse JSON data
      const data = await response.json();
    
      setUserDetails({
        ID: data.id,
        CreatedAt: new Date(data.createdAt),
        UpdatedAt: data.UpdatedAt ? new Date(data.updatedAt) : null,
        DeletedAt: data.DeletedAt ? new Date(data.deletedAt) : null,
        Email: data.email,
        Username: data.username,
        UserType: data.userType,
      });
      setIsAuthenticated(true);

      // Start the timeout when the states are set
      const timeout = window.setTimeout(() => {
        clearStates();
      }, 60 * 60 * 1000); // 1 hour in milliseconds

      setClearStatesTimeout(timeout);
  } catch (error) {
    console.log(error);
    clearStates();
  }
      
  }

  function clearStates() {
    setUserDetails(null);
    setIsAuthenticated(false);
    setClearStatesTimeout(null);
    localStorage.removeItem('token');
  }

  useEffect(() => {
    return () => {
      if (clearStatesTimeout) {
        window.clearTimeout(clearStatesTimeout);
      }
    };
  }, []); 

  return { userDetails, isAuthenticated, fetchUserDetails, login };
}