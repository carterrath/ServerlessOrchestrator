import orcaArt from '../../assets/images/orcaArt.png';

interface IProps {
  children: React.ReactNode;
}

export function BackgroundImage(props: IProps) {
  return (
    <div
      style={{
        backgroundImage: `url(${orcaArt})`,
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        height: 'calc(100svh - 80px)',
      }}
    >
      {props.children}
    </div>
  );
}
