import { useEffect } from "react";
import { IMicroserviceData } from "../../types/microservice-data";
import PlaySvg from "../../assets/svg/play.svg";
import GithubBlackSvg from "../../assets/svg/github-black.svg";
import OutputSvg from "../../assets/svg/output.svg";
interface IProps{
    item: IMicroserviceData;
}
export function MicroserviceCard(props: IProps) {
  const data=useMicroserviceCard(props);
  return (
    <div className="p-2 bg-gray-200 rounded-xl drop-shadow-lg hover:scale-[101%] transition duration-150 ease-in-out">
      <div className="flex justify-between mb-4">
        <div className="flex gap-3 justify-start items-center mb-2">
          <h2 className="text-2xl font-semibold">{props.item.FriendlyName}</h2>
          <button className="bg-green-500 rounded-lg py-1 px-2 hover:shadow-md" onClick={data.handlePlayClick}>
            <img src={PlaySvg} alt="filter" className="w-4 h-4" />
          </button>
        </div>
        <span className={`h-6 w-6 drop-shadow-lg shadow-black ${props.item.IsActive === true ? 'bg-green-500' : 'bg-red-500'} text-white rounded-full`}>
        </span>
      </div>
    
      <div className="flex justify-start items-center mb-4">
        <div className="w-8 flex justify-start items-center">
          <img src={GithubBlackSvg} alt="filter" className="w-7 h-7" />
        </div>
        <a href={props.item.RepoLink} className="text-blue-500 hover:underline">{props.item.RepoLink}</a>
      </div>
    
      <div className="flex justify-start items-center mb-4">
        <div className="w-8 flex justify-start items-center">
          <img src={OutputSvg} alt="filter" className="w-6 h-6" />
        </div>
        {props.item.OutputLink === "" ? <p className="text-gray-600">N/A</p> :
          <a href={props.item.OutputLink} className="text-blue-500 hover:underline">{props.item.OutputLink}</a>
        }
      </div>
    
      <div className="mt-4">
        <p className="text-lg">{props.item.User.Username} (Developer)</p>
        <p className="text-gray-600 text-sm">{props.item.User.Email}</p>
      </div>
    
      <div className="flex justify-between mt-4">
        <p className="text-gray-600 text-sm">Created At: {new Date(props.item.CreatedAt).toLocaleString()}</p>
        {props.item.updatedAt && (
          <p className="text-gray-600 text-sm">Updated At: {new Date(props.item.updatedAt).toLocaleString()}</p>
        )}
        {props.item.DeletedAt && (
          <p className="text-gray-600 text-sm">Deleted At: {new Date(props.item.DeletedAt).toLocaleString()}</p>
        )}
      </div>
    </div>    
  );
}

function useMicroserviceCard(props: IProps) {
  function handlePlayClick() {
    fetch('http://localhost:8080/runmicroservice', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        value: props.item.BackendName,
      }),
    })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  
  return { handlePlayClick };
}

