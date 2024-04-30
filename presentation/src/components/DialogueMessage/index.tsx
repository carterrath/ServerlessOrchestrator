import { useEffect, useState, useRef } from 'react';
import closeSvg from '../../assets/svg/close.svg';

interface IProps {
  title: string;
  message: string;
  type: 'success' | 'error';
  show: boolean;
  onClose: () => void;
}

export function DialogueMessage(props: IProps) {
  // Define color schemes based on the type of message
  const bgColor = props.type === 'success' ? 'bg-green-500' : 'bg-red-500';
  const buttonColor = props.type === 'success' ? 'bg-green-800' : 'bg-red-800';

  const [show, setShow] = useState(false);
  const dialogueRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    setShow(props.show);
  }, [props.show]);

  useEffect(() => {
    function handleOutsideClick(e: MouseEvent) {
      if (dialogueRef.current && !dialogueRef.current.contains(e.target as Node)) {
        setShow(false);
        props.onClose();
      }
    }

    document.addEventListener('mousedown', handleOutsideClick);

    return () => {
      document.removeEventListener('mousedown', handleOutsideClick);
    };
  }, []);

  return (
    show && (
      <div className="fixed inset-0 flex items-end justify-center p-4 z-50">
        <div
          className={`w-full max-w-md rounded-lg shadow-lg p-6 transform translate-y-full animate-slide-up ${bgColor}`}
          ref={dialogueRef}
        >
          <div className="flex justify-between items-center">
            <h2 className="text-lg font-bold text-white">{props.title}</h2>
            <button
              onClick={() => {
                setShow(false);
                props.onClose();
              }}
              className={`p-2 rounded-full hover:shadow-md ${buttonColor}`}
            >
              <img src={closeSvg} alt="close" className="w-4 h-4" />
            </button>
          </div>
          <p className="text-white my-2">{props.message}</p>
        </div>
      </div>
    )
  );
}
