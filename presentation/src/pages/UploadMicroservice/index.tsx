import React, { useState } from 'react';
import UploadSvg from "../../assets/svg/upload.svg";
import PlusSvg from "../../assets/svg/plus.svg";
import MinusSvg from "../../assets/svg/minus.svg";




export function UploadMicroservice() {
  const data = useUploadMicroservice();

  return (
    <>
    <div className=" items-center">
      <div className="p-2 bg-gray-100 my-12 mx-32 rounded-xl drop-shadow-lg">
        <form onSubmit={data.handleSubmit}>
          <div className="flex items-center mx-4">
            <div className="font-extrabold m-4 text-2xl">
                  Upload your Function!
            </div>
            <button className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md">
                  <img src={UploadSvg} alt="upload" className="w-8 h-8"/>
            </button>
          </div>
        <div className="flex items-center mx-4">
            <div className="w-1/2">
              <div className="my-2 text-sm">
                  Microservice Name
              </div>
              <input
                  type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
            </div>
            <div className="w-1/2">
              <div className="my-2 text-sm">
                  Github Repository Link
              </div>
              <input
                  type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>            
            </div>
        </div>
        <div className="flex items-center mx-4">
            <div className="w-1/2">
              <div className="my-2 text-sm">
                  Input Name
              </div>
              <input
                  type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
            </div>
            <div className="w-1/2">
              <div className="my-2 text-sm">
                  Input Data Type
              </div>
              <input
                  type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>   
              <button className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md">
                  <img src={PlusSvg} alt="add" className="w-8 h-8"/>
              </button>  
              <button className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md">
                  <img src={MinusSvg} alt="add" className="w-8 h-8"/>
              </button>         
            </div>
        </div>
        <div className="flex items-center mx-4">
            <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md">
                  Add Input
            </button>
        </div>
        </form>
      </div>
    </div>
    </>
  );
}

function useUploadMicroservice(){
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