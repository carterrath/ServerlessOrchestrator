import React, { useEffect, useState } from 'react';
import UploadSvg from '../../assets/svg/upload.svg';
import PlusSvg from '../../assets/svg/plus.svg';
import MinusSvg from '../../assets/svg/minus.svg';
import { IMicroserviceUpload } from '../../types/microservice-upload';
import { useAuth } from '../../hooks/useAuth';

export function UploadMicroservice() {
  const data = useUploadMicroservice();

  return (
    <div className="flex justify-center items-center py-8">
      <div className="p-4 bg-gray-200 w-2/3 rounded-xl drop-shadow-lg">
        <form onSubmit={data.handleSubmit}>
          <div className="flex items-center justify-between mb-4">
            <div className="flex">
              <div className="font-extrabold text-2xl">Upload your Microservice!</div>
              {data.resultMessage && (
                <div
                  className={`
                    ${data.resultMessage.type === 'success' && 'text-green-600'} 
                    ${data.resultMessage.type === 'error' && 'text-red-600'}
                  `}
                >
                  {data.resultMessage.msg}
                </div>
              )}
            </div>
            <button className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md">
              <img src={UploadSvg} alt="upload" className="w-8 h-8" />
            </button>
          </div>
          <div className="flex items-center gap-2 mb-2">
            <div className="w-1/2">
              <div className="text-sm mb-2">Microservice Name</div>
              <input
                type="text"
                name="FriendlyName"
                placeholder="My Microservice"
                value={data.microservice.FriendlyName}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-full border-gray-300 hover:shadow-md"
                required
              />
            </div>
            <div className="w-1/2">
              <div className="text-sm mb-2">Github Repository Link</div>
              <input
                type="text"
                name="RepoLink"
                placeholder="https://github.com/janedoe/MyRepository.git"
                value={data.microservice.RepoLink}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-full border-gray-300 hover:shadow-md"
                required
              />
            </div>
          </div>
          {
            // data.microservice.Inputs.length === 0 && (
            //   <div className="flex items-center justify-center">
            //     <button onClick={data.handleAddInput} className="bg-gray-300 rounded-lg mt-4 py-2 px-2 hover:shadow-md">
            //       Add Input
            //     </button>
            //   </div>
            // )
          }
          {data.microservice.Inputs.map((input, index) => (
            <div className="flex items-center gap-2" key={index}>
              <div className="w-1/2">
                <div className="my-2 text-sm">Input Name</div>
                <input
                  type="text"
                  name={`Inputs[${index}].Name`}
                  value={input.Name}
                  onChange={data.handleInputChange(index)}
                  className="rounded-lg p-2 border w-full border-gray-300 hover:shadow-md"
                  required
                />
              </div>
              <div className="w-1/2">
                <div className="my-2 text-sm">Input Data Type</div>
                <div className="flex items-center w-full justify-between">
                  <input
                    type="text"
                    name={`Inputs[${index}].DataType`}
                    value={input.DataType}
                    onChange={data.handleInputChange(index)}
                    className="rounded-lg p-2 border border-gray-300 hover:shadow-md"
                    required
                  />
                  <div className="flex gap-3">
                    {index === data.microservice.Inputs.length - 1 && (
                      <button
                        type="button"
                        onClick={data.handleAddInput}
                        className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md"
                      >
                        <img src={PlusSvg} alt="add" className="w-6 h-6" />
                      </button>
                    )}
                    <button
                      type="button"
                      onClick={() => data.handleRemoveInput(index)}
                      className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md"
                    >
                      <img src={MinusSvg} alt="remove" className="w-6 h-6" />
                    </button>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </form>
      </div>
    </div>
  );
}

interface IResultMessage {
  msg: string;
  type: 'error' | 'success';
}

function useUploadMicroservice() {
  const [resultMessage, setResultMessage] = useState<IResultMessage | null>(null);
  const [microservice, setMicroservice] = useState<IMicroserviceUpload>({
    FriendlyName: '',
    RepoLink: '',
    UserID: 0,
    Inputs: [],
  });

  const auth = useAuth();

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target;
    setMicroservice((prevMicroservice) => ({
      ...prevMicroservice,
      [name]: value,
    }));
  }

  function handleInputChange(index: number) {
    return (e: React.ChangeEvent<HTMLInputElement>) => {
      const { name, value } = e.target;
      setMicroservice((prevMicroservice) => ({
        ...prevMicroservice,
        Inputs: prevMicroservice.Inputs.map((input, i) =>
          i === index ? { ...input, [name.split('.')[1]]: value } : input,
        ),
      }));
    };
  }

  function handleAddInput() {
    setMicroservice((prevMicroservice) => ({
      ...prevMicroservice,
      Inputs: [...prevMicroservice.Inputs, { MicroserviceID: '', Id: 0, Name: '', DataType: '' }],
    }));
  }

  function handleRemoveInput(index: number) {
    setMicroservice((prevMicroservice) => ({
      ...prevMicroservice,
      Inputs: prevMicroservice.Inputs.filter((input, i) => i !== index),
    }));
  }

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    console.log(microservice);
    // Convert the microservice object to JSON
    const microserviceJson = JSON.stringify(microservice);

    // Make a POST request to the endpoint
    fetch('http://localhost:8080/microservice', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: microserviceJson,
    })
      .then((response) => {
        if (!response.ok) {
          // Handle error response
          return response.json().then((error) => {
            throw new Error(error.error);
          });
        }
        return response.json(); // If response is OK, parse JSON data
      })
      .then((res) => {
        console.log('Microservice uploaded successfully:', res.message);
        setResultMessage({ msg: res.message, type: 'success' });
      })
      .catch((error) => {
        console.log(error);
        setResultMessage({ msg: error.message, type: 'error' });
      });
  }

  useEffect(() => {
    auth?.fetchUserDetails();
  }, []);

  useEffect(() => {
    setMicroservice((prevMicroservice) => ({
      ...prevMicroservice,
      UserID: auth?.userDetails?.ID || 0,
    }));
  }, [auth?.userDetails]);

  return {
    handleSubmit,
    handleChange,
    handleInputChange,
    handleAddInput,
    handleRemoveInput,
    microservice,
    resultMessage,
  };
}
