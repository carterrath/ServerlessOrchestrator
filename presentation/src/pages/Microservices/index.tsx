import React, { useEffect, useState } from "react";
import { ScrollableList } from "../../components/ScrollableList";

export function Microservices() {
    const data = useMicroservices();
    const microserviceNames = data.microservices?.map(microservice => microservice.Name);
    return (
        <div>
            {microserviceNames == null ? <p>Loading</p>:
             <ScrollableList items={microserviceNames} />
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

interface IMicroservice {
    ID: number;
    Name: string;
    ServiceHook: string;
    BuildScript: string;
    PlaceHolder: string;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
  }