import { IUserData } from '../types/user-data';
import { useState, useEffect, createContext, useMemo, useContext } from 'react';
import { useNavigate } from 'react-router-dom';

type UserType = 'Developer' | 'Consumer';

interface IAuthContext {
  userDetails: IUserData | null;
  isAuthenticated: boolean;
  fetchUserDetails: () => void;
  login: (username: string, password: string, userType: UserType) => Promise<string>;
  logout: () => void;
}

export const AuthContext = createContext<IAuthContext | null>(null);

interface IAuthProviderProps {
  children: React.ReactNode;
}

export function AuthProvider(props: IAuthProviderProps) {
  const [userDetails, setUserDetails] = useState<IUserData | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [clearStatesTimeout, setClearStatesTimeout] = useState<number | null>(null);

  const navigate = useNavigate();

  async function login(username: string, password: string, userType: UserType) {
    try {
      if (userType === 'Developer') {
        const response = await fetch('https://serverlessorchestrator.com/login/developer', {
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
          fetchUserDetails();
          return 'success';
        } else {
          const errorData = await response.json();
          return errorData.error;
        }
      }
      if (userType === 'Consumer') {
        const response = await fetch('https://serverlessorchestrator.com/login/consumer', {
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
          fetchUserDetails();
          return 'success';
        } else {
          const errorData = await response.json();
          return errorData.error;
        }
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
    if (!token) {
      // maybe redirect to login page
      return;
    }
    try {
      const response = await fetch('https://serverlessorchestrator.com/getuserdetails', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`, // Add the token to the Authorization header
        },
        credentials: 'include',
      });

      if (!response.ok) {
        throw new Error('Failed to fetch user details');
      }
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
      const timeout = window.setTimeout(
        () => {
          clearStates();
        },
        60 * 60 * 1000,
      ); // 1 hour in milliseconds

      setClearStatesTimeout(timeout);
    } catch (error) {
      clearStates();
    }
  }

  function clearStates() {
    setUserDetails(null);
    setIsAuthenticated(false);
    if (clearStatesTimeout) {
      window.clearTimeout(clearStatesTimeout);
    }
    setClearStatesTimeout(null);
    localStorage.removeItem('token'); // This ensures the token is removed on logout
  }

  function logout() {
    clearStates(); // Reuse clearStates to handle the cleanup
    navigate('/Home');
    // Optionally, navigate the user to the login page or another page as appropriate
    // navigate('/login'); Uncomment and adjust according to your routing setup
  }

  useEffect(() => {
    return () => {
      if (clearStatesTimeout) {
        window.clearTimeout(clearStatesTimeout);
      }
    };
  }, [clearStatesTimeout]);

  //return { userDetails, isAuthenticated, fetchUserDetails, login, logout };

  const contextValue = useMemo(
    () => ({
      userDetails,
      isAuthenticated,
      fetchUserDetails,
      login,
      logout,
    }),
    [userDetails, isAuthenticated, fetchUserDetails, login, logout],
  );

  return <AuthContext.Provider value={contextValue}>{props.children}</AuthContext.Provider>;
}

export function useAuth() {
  return useContext<IAuthContext | null>(AuthContext);
}
