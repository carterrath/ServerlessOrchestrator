// MainChoice.tsx
import { useNavigate } from 'react-router-dom';

function MainChoice() {
  const navigate = useNavigate();

  const handleDeveloperSelect = () => {
    navigate('/developer-signup');
  };

  const handleConsumerSelect = () => {
    navigate('/consumer-signup');
  };

  return (
    <div className="container mx-auto p-4 lg:w-2/5">
      <div className="p-2 bg-gray-200 my-12 mx-auto rounded-xl drop-shadow-lg">
      <div className="flex items-center justify-center font-extrabold m-4 text-2xl"> Select your user type! </div>
        <div className="flex items-center justify-center mb-4">
          <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md" onClick={handleDeveloperSelect}>Developer</button>
          <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md" onClick={handleConsumerSelect}>Consumer</button>
        </div>
      </div>
    </div>
  );
}

export default MainChoice;
