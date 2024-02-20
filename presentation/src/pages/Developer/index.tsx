import React, { useState } from 'react';


export function Developer() {
  const data = useDeveloper();

  return (
    <div className="developer-page">
      <header className="developer-page-header">
        <h1>Serverless Orchestrator</h1>
        <h2>Connect your app!</h2>
        <form className="u-p-t-5" onSubmit={data.handleSubmit}>
          <input
            className="u-w-quarter u-m-r-5 u-p-1"
            type="text"
            placeholder="Enter Function Name"
            value={data.functionName}
            onChange={(event) => data.setFunctionName(event.target.value)}
          />
          <input
            className="u-w-quarter u-m-r-5 u-p-1"
            type="text"
            placeholder="Enter GitHub Repository Link"
            value={data.repoLink}
            onChange={(event) => data.setRepoLink(event.target.value)}
          />
          <input
            className="u-w-quarter u-m-r-5 u-p-1"
            type="text"
            placeholder="Enter User Input"
            value={data.userInput}
            onChange={(event) => data.setUserInput(event.target.value)}
          />
          <button className="u-w-tenth u-p-1" type="submit">Connect</button>
        </form>
      </header>
    </div>
  );
}

function useDeveloper(){
  const [functionName, setFunctionName] = useState('');
  const [repoLink, setRepoLink] = useState('');
  const [userInput, setUserInput] = useState('');

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    
    const data = {
      name: functionName,
      placeholder: repoLink,
      input: userInput
    };

    try {
      const response = await fetch('http://localhost:8080/microservice', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      });

      if (!response.ok) {
        throw new Error('Failed to submit data');
      }

      // If the request was successful, reset the form inputs
      setFunctionName('');
      setRepoLink('');
      setUserInput('');

      console.log('Data submitted successfully');
    } catch (error) {
      console.error('Error submitting data:', error);
    }
  };
  
  return {
    functionName,
    setFunctionName,
    repoLink,
    setRepoLink,
    userInput,
    setUserInput,
    handleSubmit
  }

}