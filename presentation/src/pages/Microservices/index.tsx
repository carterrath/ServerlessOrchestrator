import { useEffect, useState } from 'react';
import { IMicroserviceData } from '../../types/microservice-data';
import { MicroserviceCards } from './MicroserviceCards';
import FilterSvg from '../../assets/svg/filter.svg';
import UploadSvg from '../../assets/svg/upload.svg';
import { BackgroundGradient } from '../../components/BackgroundGradient';

export function Microservices() {
  const data = useMicroservices();
  return (
    <BackgroundGradient>
      <div>
        {
          <>
            <div className="flex justify-between mx-8 items-center">
              <div className="font-semibold my-4 text-2xl">
                Microservices ({data.microservices !== null ? data.microservices.length : 0})
              </div>
              <div className="flex items-center gap-2 w-1/3">
                <button
                  className="bg-gray-800 rounded-lg py-1 px-2 hover:shadow-md"
                  onClick={() => data.handleUploadClick()}
                >
                  <img src={UploadSvg} alt="upload" className="w-8 h-8" />
                </button>
                <button className="bg-gray-800 rounded-lg py-1 px-2 hover:shadow-md">
                  <img src={FilterSvg} alt="filter" className="w-8 h-8" />
                </button>
                <input
                  type="text"
                  placeholder="Search"
                  className="rounded-lg p-2 border w-full border-gray-300 hover:shadow-md"
                  value={data.search}
                  onChange={(e) => data.setSearch(e.target.value)}
                />
              </div>
            </div>
            {data.microservices !== null && data.microservices.length > 0 && (
              <MicroserviceCards items={data.microservices} search={data.search} />
            )}
          </>
        }
      </div>
    </BackgroundGradient>
  );
}

function useMicroservices() {
  const [microservices, setMicroservices] = useState<IMicroserviceData[]>([]);
  const [search, setSearch] = useState<string>('');

  function handleUploadClick() {
    window.location.href = '/UploadMicroservice';
  }

  useEffect(() => {
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

    // Call getMicroservices immediately and then every 5 seconds
    getMicroservices();
    const intervalId = setInterval(getMicroservices, 5000);

    // Clean up the interval on unmount
    return () => {
      clearInterval(intervalId);
    };
  }, []);

  return {
    microservices,
    handleUploadClick,
    search,
    setSearch,
  };
}
