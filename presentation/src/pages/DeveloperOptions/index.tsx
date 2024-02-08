
import './../../assets/styles/DeveloperOptions.css';
import { Routes, Route, useNavigate } from 'react-router-dom';
//import DeveloperPage from './DeveloperPage';

export default function DeveloperOptions() {
    const data = useDeveloperOptions();
    const navigate = useNavigate();
    const handleUpload = () => {
        navigate('/developer');
    };
    
    return (
        <div className="App">
        <Routes>
        <Route path="/" element={
            // This is your home page layout with login options
            <header className="App-header">
            <h1>Serverless Orchestrator</h1>
            <div className="login-container">
            <div className="login developer-login">
            <button onClick={handleUpload}>Upload Microservice</button>
            </div>
            <div className="login user-login">
            <div className="dropdown">
            <button onClick={data.getMicroservices}>Browse Microservices</button>
            <button className="dropdown-btn">Browse Microservices</button>
            <div className="dropdown-content">
            <a href="#1">Cart</a>
            <a href="#2">Catalog</a>
            <a href="#3">Payment</a>
            <a href="#4">Shipping</a>
            <a href="#5">Label</a>
            </div>
            </div>
            </div>
            </div>
            </header>
        } />
        </Routes>
        </div>
        );
    }
    
   export function useDeveloperOptions() {

    const getMicroservices = async () => {
        try {
          const response = await fetch('http://localhost:8080/microservice');
          if (!response.ok) {
            throw new Error('Failed to fetch microservices');
          }
          const data = await response.json();
          console.log(data);
        } catch (error) {
          console.error('Error fetching microservices:', error);
        }
      };

    return {
        getMicroservices
    }
   }
    