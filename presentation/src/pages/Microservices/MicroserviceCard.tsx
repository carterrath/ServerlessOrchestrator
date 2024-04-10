import { IMicroserviceData } from "../../types/microservice-data";
import { format } from 'date-fns';
import PlaySvg from "../../assets/svg/play.svg";
import StopSvg from "../../assets/svg/stop.svg";
import GithubBlackSvg from "../../assets/svg/github-black.svg";
import OutputSvg from "../../assets/svg/output.svg";
interface IProps {
  item: IMicroserviceData;
}
export function MicroserviceCard(props: IProps) {
  const data = useMicroserviceCard(props);
  return (
    <div className="p-2 bg-gray-200 rounded-xl drop-shadow-lg hover:scale-[101%] transition duration-150 ease-in-out">
      <div className="flex justify-between mb-4">
        <div className="flex gap-3 justify-start items-center mb-2">
          <h2 className="text-2xl font-semibold">{props.item.FriendlyName}</h2>
          <button className={`${props.item.IsActive === true ? 'bg-red-500': 'bg-green-500'} rounded-lg py-1 px-2 hover:shadow-md`} onClick={data.handlePlayClick}>
            <img src={props.item.IsActive === true? StopSvg: PlaySvg} alt="filter" className="w-4 h-4" />
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
        <p className="text-gray-600 text-sm">Created At: {data.formatDate(props.item.CreatedAt)}</p>
        {props.item.UpdatedAt && (
          <p className="text-gray-600 text-sm">Updated At: {data.formatDate(props.item.UpdatedAt)}</p>
        )}
        {props.item.DeletedAt && props.item.DeletedAt !== "0001-01-01 00:00:00 +0000 UTC" && (
          <p className="text-gray-600 text-sm">Deleted At: {data.formatDate(props.item.DeletedAt)}</p>
        )}
      </div>
    </div>
  );
}

function useMicroserviceCard(props: IProps) {
  function formatDate(dateString: string) {
    // Assuming the date string is in ISO 8601 format
    const date = new Date(dateString);
    return format(date, 'MMMM dd, yyyy HH:mm:ss');
  }
  function handlePlayClick() {
    if(props.item.IsActive === false){
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
          // Update the IsActive property of the microservice
        props.item.IsActive = data.isRunning;
        })
        .catch((error) => {
          console.error('Error:', error);
        });
    }
    else{
      fetch('http://localhost:8080/stopmicroservice', {
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
          // Update the IsActive property of the microservice
        props.item.IsActive = data.isRunning;
        })
        .catch((error) => {
          console.error('Error:', error);
        });
    }
  }

  return {
    handlePlayClick,
    formatDate
  };
}

