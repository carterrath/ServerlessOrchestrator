import React, { useState } from 'react';
import UploadSvg from "../../assets/svg/upload.svg";
import PlusSvg from "../../assets/svg/plus.svg";
import MinusSvg from "../../assets/svg/minus.svg";
import { IMicroservice } from '../../types/microservice';

export function UploadMicroservice() {
  const data = useUploadMicroservice();

  return (
    <div className="items-center">
      <div className="p-2 bg-gray-200 my-12 mx-32 rounded-xl drop-shadow-lg">
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
                name="Name"
                value={data.microservice.Name}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
            <div className="w-1/2">
              <div className="my-2 text-sm">
                Github Repository Link
              </div>
              <input
                type="text"
                name="RepoLink"
                value={data.microservice.RepoLink}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
          </div>
          {
          // data.microservice.Inputs.length === 0 && (
          //   <div className="flex items-center mx-4">
          //     <button onClick={data.handleAddInput} className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md">
          //       Add Input
          //     </button>
          //   </div>
          // )
          }
          {data.microservice.Inputs.map((input, index) => (
            <div className="flex items-center mx-4" key={index}>
              <div className="w-1/2">
                <div className="my-2 text-sm">
                  Input Name
                </div>
                <input
                  type="text"
                  name={`Inputs[${index}].Name`}
                  value={input.Name}
                  onChange={data.handleInputChange(index)}
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
                />
              </div>
              <div className="w-1/2">
                <div className="my-2 text-sm">
                  Input Data Type
                </div>
                <input
                  type="text"
                  name={`Inputs[${index}].DataType`}
                  value={input.DataType}
                  onChange={data.handleInputChange(index)}
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
                />
                <button type="button" onClick={() => data.handleRemoveInput(index)} className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md">
                  <img src={MinusSvg} alt="remove" className="w-8 h-8"/>
                </button>
                {index === data.microservice.Inputs.length - 1 && (
                  <button type="button" onClick={data.handleAddInput} className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md">
                    <img src={PlusSvg} alt="add" className="w-8 h-8"/>
                  </button>
                )}
              </div>
            </div>
          ))}
        </form>
      </div>
    </div>
  );
}

function useUploadMicroservice(){
  const [microservice, setMicroservice] = useState<IMicroservice>({
    ID: 0,
    CreatedAt: new Date(),
    UpdatedAt: new Date(),
    DeletedAt: null,
    Name: '',
    RepoLink: '',
    Author: '',
    Inputs: [],
    Status: '',
  });

  function handleChange (e: React.ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target;
    setMicroservice(prevMicroservice => ({
      ...prevMicroservice,
      [name]: value
    }));
  };

  function handleInputChange(index: number) {
    return (e: React.ChangeEvent<HTMLInputElement>) => {
      const { name, value } = e.target;
      setMicroservice(prevMicroservice => ({
        ...prevMicroservice,
        Inputs: prevMicroservice.Inputs.map((input, i) =>
          i === index ? { ...input, [name.split('.')[1]]: value } : input
        )
      }));
    };
  }

  function handleAddInput() {
    setMicroservice(prevMicroservice => ({
      ...prevMicroservice,
      Inputs: [...prevMicroservice.Inputs, { MicroserviceID: '', Id: 0, Name: '', DataType: '' }]
    }));
  }

  function handleRemoveInput(index: number) {
    setMicroservice(prevMicroservice => ({
      ...prevMicroservice,
      Inputs: prevMicroservice.Inputs.filter((input, i) => i !== index)
    }));
  }

  function handleSubmit (e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    console.log(microservice);
    // Convert the microservice object to JSON
    const microserviceJson = JSON.stringify(microservice);

    // Make a POST request to the endpoint
    fetch('http://localhost:8080/microservice', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: microserviceJson
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to upload microservice');
        }
        return response.json();
    })
    .then(data => {
        console.log('Microservice uploaded successfully:', data);
        // Optionally, you can reset the form or perform any other actions here
    })
    .catch(error => {
        console.error('Error uploading microservice:', error);
        // Optionally, you can handle the error and display a message to the user
    });
  };

  return {
    handleSubmit,
    handleChange,
    handleInputChange,
    handleAddInput,
    handleRemoveInput,
    microservice
  }
}
