import React from "react";
import { useState } from "react";

export function ScrollAssistant() {
    const [isTop, setIsTop] = useState(true);

    const handleScroll = () => {
        const position = window.pageYOffset;
        setIsTop(position < 50);
    };

    React.useEffect(() => {
        window.addEventListener('scroll', handleScroll, { passive: true });

        return () => {
            window.removeEventListener('scroll', handleScroll);
        };
    }, []);

    return (
        <button
            onClick={() => window.scrollTo({ top: isTop ? document.body.scrollHeight : 0, behavior: 'smooth' })}
            className="fixed bottom-4 left-1/2 transform -translate-x-1/2 bg-black text-white w-12 h-12 rounded-full flex items-center justify-center text-2xl animate-bounce200"
        >
            {isTop ? '↓' : '↑'}
        </button>
    );
}