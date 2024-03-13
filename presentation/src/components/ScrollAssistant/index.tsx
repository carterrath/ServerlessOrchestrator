import React, { useState, useEffect } from "react";

export function ScrollAssistant() {
    const sectionIds = ['header-section', 'features-section', 'about-section', 'other-materials-section', 'team-section'];
    const [currentSectionIndex, setCurrentSectionIndex] = useState(0);

    const scrollToSection = (index: number) => {
        const section = document.getElementById(sectionIds[index]);
        if (section) {
            window.scrollTo({ top: section.offsetTop, behavior: 'smooth' });
        }
    };

    const handleScroll = () => {
        const position = window.pageYOffset;
        const currentSection = sectionIds.findIndex((id, index) => {
            const section = document.getElementById(id);
            return section && section.offsetTop >= position; // Add null check for section
        });
        setCurrentSectionIndex(currentSection >= 0 ? currentSection : sectionIds.length - 1);
    };

    useEffect(() => {
        window.addEventListener('scroll', handleScroll, { passive: true });
        return () => {
            window.removeEventListener('scroll', handleScroll);
        };
    }, []);

    return (
        <button
            onClick={() => scrollToSection((currentSectionIndex + 1) % sectionIds.length)}
            className="fixed bottom-4 left-1/2 transform -translate-x-1/2 bg-black text-white w-12 h-12 rounded-full flex items-center justify-center text-2xl"
        >
            ↓
        </button>
    );
}
