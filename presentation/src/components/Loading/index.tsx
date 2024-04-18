import loadingWhiteSvg from '../../assets/svg/loading-white.svg';
import loadingBlackSvg from '../../assets/svg/loading-black.svg';

interface IProps {
  color: 'white' | 'black';
  sm?: boolean;
}
export function Loading(props: IProps) {
  return (
    <img
      src={props.color === 'white' ? loadingWhiteSvg : loadingBlackSvg}
      alt="loading"
      className={`animate-spin ${props.sm ? 'w-4 h-4' : 'w-8 h-8'}`}
    />
  );
}
