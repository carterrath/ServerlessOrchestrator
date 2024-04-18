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
        <Route
          path="/"
          element={
            // This is your home page layout with login options
            <header className="App-header">
              <h1>Serverless Orchestrator</h1>
              <div className="login-container">
                <div className="login DeveloperLogin">
                  <button onClick={handleUpload}>Upload Microservice</button>
                </div>
                <div className="login user-login">
                  <div className="dropdown">
                    <button onClick={data.handleBrowseMicroservices}>Browse Microservices</button>
                  </div>
                </div>
              </div>
            </header>
          }
        />
      </Routes>
    </div>
  );
}

export function useDeveloperOptions() {
  function handleBrowseMicroservices() {
    window.location.href = '/Microservices';
  }

  return {
    handleBrowseMicroservices,
  };
}
