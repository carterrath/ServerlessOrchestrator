import React, { useEffect, useState } from "react";
import { ScrollableList } from "../../components/ScrollableList";
import { IMicroservice } from "../../types/microservice";


export function Microservices() {
    const data = useMicroservices();
    const microserviceNames = data.microservices?.map(microservice => microservice.Name);
    return (
        <div>
            {microserviceNames == null ? <p>Loading</p>:
            <>
            <h1>Microservices</h1>
             <ScrollableList items={microserviceNames} />
             </>
            }
        </div>
    );
}

function useMicroservices() {

    const [microservices, setMicroservices] = useState<IMicroservice[]>([]);

    const getMicroservices = async () => {
        try {
          const response = await fetch('http://localhost:8080/microservice');
          if (!response.ok) {
            throw new Error('Failed to fetch microservices');
          }
          const data = await response.json();
          setMicroservices(data);
        } catch (error) {
          console.error('Error fetching microservices:', error);
        }
      };

    useEffect(() => {
        getMicroservices();
    }, []);

  return {
    microservices,
  };
}

