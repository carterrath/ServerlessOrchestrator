import React, { useState } from 'react';



export function Developer() {
  const data = useDeveloper();

  return (
    <form onSubmit={data.handleSubmit} className="max-w-3xl mx-auto shadow-lg p-6 bg-white">
      <h2 className="text-base font-semibold leading-7 text-gray-900">Connect your App!</h2>
        <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
            <div className="sm:col-span-3">
              <label htmlFor="first-name" className="block text-sm font-medium leading-6 text-gray-900">
                Microservice Name
              </label>
              <div className="mt-2">
                <input
                  type="text"
                  placeholder="  Enter Microservice Name"
                  value={data.functionName}
                  onChange={(event) => data.setFunctionName(event.target.value)}
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div className="sm:col-span-3">
              <label htmlFor="last-name" className="block text-sm font-medium leading-6 text-gray-900">
                Github Repository Link
              </label>
              <div className="mt-2">
                <input
                  type="text"
                  placeholder="  Enter GitHub Repository Link"
                  value={data.repoLink}
                  onChange={(event) => data.setRepoLink(event.target.value)}
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div className="col-span-full">
              <label htmlFor="street-address" className="block text-sm font-medium leading-6 text-gray-900">
                Input
              </label>
              <div className="mt-2">
                <input
                  type="text"
                  placeholder="  Enter User Input"
                  value={data.userInput}
                  onChange={(event) => data.setUserInput(event.target.value)}
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
          </div>

          <div className="mt-6 flex items-center justify-end gap-x-6">
            <button
              type="submit"
              className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              Connect
            </button>
          </div> 
    </form>
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