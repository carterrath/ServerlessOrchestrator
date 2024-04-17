interface IProps {
  children: React.ReactNode;
}

export function BackgroundGradient(props: IProps) {
  return (
    <div
      className="bg-gradient-to-t from-amaranthPink to-white"
      style={{ backgroundSize: 'cover', backgroundPosition: 'center', height: 'calc(100svh - 80px)' }}
    >
      {props.children}
    </div>
  );
}
