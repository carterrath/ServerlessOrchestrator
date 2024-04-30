import { useEffect, useState } from 'react';
import { IMicroserviceData } from '../../types/microservice-data';
import { MicroserviceCards } from './MicroserviceCards';
import UploadSvg from '../../assets/svg/upload.svg';
import { BackgroundGradient } from '../../components/BackgroundGradient';
import { API_URL } from '../../constants';

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
              <MicroserviceCards
                items={data.microservices}
                search={data.search}
                getMicroservices={data.getMicroservices}
              />
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

  function getMicroservices() {
    fetch(`${API_URL}/api/microservice`)
      .then((response) => {
        if (!response.ok) {
          throw new Error('Failed to fetch microservices');
        }
        return response.json();
      })
      .then((data) => {
        setMicroservices(data);
      })
      .catch((error) => {
        console.error('Error fetching microservices:', error);
      });
  }

  useEffect(() => {
    getMicroservices();
  }, []);

  return {
    microservices,
    handleUploadClick,
    search,
    setSearch,
    getMicroservices,
  };
}
