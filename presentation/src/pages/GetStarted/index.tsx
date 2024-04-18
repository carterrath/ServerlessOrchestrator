// MainChoice.tsx
import { useNavigate } from 'react-router-dom';
import toolsSvg from '../../assets/svg/tools.svg';
import shoppingBagSvg from '../../assets/svg/shopping-bag.svg';
import { BackgroundGradient } from '../../components/BackgroundGradient';

export default function MainChoice() {
  const navigate = useNavigate();

  const handleDeveloperSelect = () => {
    navigate('/developer-signup');
  };

  const handleConsumerSelect = () => {
    navigate('/consumer-signup');
  };

  return (
    <BackgroundGradient>
      <div className="container mx-auto p-4">
        <div className="flex items-center justify-center font-extrabold m-4 text-4xl"> Select user type </div>
        <div className="flex gap-6 justify-center">
          <div className="flex flex-col items-center p-2 bg-gray-200 my-6 rounded-xl drop-shadow-lg w-1/2">
            <div className="flex items-center justify-center font-extrabold m-4 text-2xl"> Developer </div>
            {/* image of tools */}
            <img src={toolsSvg} alt="developer" className="rounded-lg w-20 h-20 pb-6" />
            <div className="flex flex-col justify-between">
              <div>
                Developers are the creative backbone of the application, responsible for uploading and managing their
                microservices on the global hub, ensuring they are functional, well-documented, and meet quality
                standards. They handle coding, testing, and version control.
              </div>
              <div className="flex items-center justify-center mb-4">
                <button
                  className="bg-gray-800 text-white rounded-lg m-4 py-2 px-2 hover:shadow-md w-full"
                  onClick={handleDeveloperSelect}
                >
                  Select Developer
                </button>
              </div>
            </div>
          </div>

          <div className="flex flex-col items-center p-2 bg-gray-200 my-6 rounded-xl drop-shadow-lg w-1/2">
            <div className="flex items-center justify-center font-extrabold m-4 text-2xl"> Consumer </div>
            {/* image of shopping bag */}
            <img src={shoppingBagSvg} alt="consumer" className="rounded-lg w-20 h-20 pb-6" />
            <div className="flex flex-col justify-between">
              <div>
                Consumers are the end-users who utilize the microservices available on the global hub to enhance their
                own applications and systems. They do not contribute code; instead, they browse the hub to find
                microservices that suit their needs, deploy them, and integrate them into their projects.
              </div>
              <div className="flex items-center justify-center mb-4">
                <button
                  className="bg-gray-800 text-white rounded-lg m-4 py-2 px-2 hover:shadow-md w-full"
                  onClick={handleConsumerSelect}
                >
                  Select Consumer
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </BackgroundGradient>
  );
}
