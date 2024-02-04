
import './../../assets/styles/DeveloperOptions.css';
import { Routes, Route, useNavigate } from 'react-router-dom';
//import DeveloperPage from './DeveloperPage';

function DeveloperOptions() {
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
    
    export default DeveloperOptions;
    