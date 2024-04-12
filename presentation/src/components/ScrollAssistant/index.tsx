import React, { useState, useEffect } from 'react';

export function ScrollAssistant() {
  const sectionIds = ['header-section', 'features-section', 'about-section', 'other-materials-section', 'team-section'];
  const [currentSectionIndex, setCurrentSectionIndex] = useState(0);
  const [isAtBottom, setIsAtBottom] = useState(false);

  const scrollToSection = (index: number) => {
    // Get the section element
    const section = document.getElementById(sectionIds[index]);
    if (section) {
      // Scroll to the section
      window.scrollTo({ top: section.offsetTop, behavior: 'smooth' });
    }
  };

  // Function to handle scroll events
  const handleScroll = () => {
    // Get the current scroll position
    const position = window.scrollY;

    // Find the current section based on the scroll position
    const currentSection = sectionIds.findIndex((id) => {
      const section = document.getElementById(id);
      // Check if the section's top offset is greater than or equal to the scroll position
      return section && section.offsetTop >= position;
    });

    // Update the current section index state
    setCurrentSectionIndex(currentSection >= 0 ? currentSection : sectionIds.length - 1);

    // Check if the page is at the bottom
    const atBottom = window.innerHeight + window.scrollY >= document.documentElement.scrollHeight;
    setIsAtBottom(atBottom);
  };

  // Use effect hook to add and remove the scroll event listener
  useEffect(() => {
    // Add the scroll event listener when the component mounts
    window.addEventListener('scroll', handleScroll, { passive: true });

    // Remove the scroll event listener when the component unmounts
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, []); // Empty dependency array means this effect runs once on mount and cleanup on unmount

  // Render a button that scrolls to the next section when clicked
  return (
    <button
      onClick={() => scrollToSection((currentSectionIndex + 1) % sectionIds.length)}
      className="fixed bottom-4 left-1/2 transform -translate-x-1/2 bg-black text-white w-12 h-12 rounded-full flex items-center justify-center text-2xl animate-bounce"
    >
      {isAtBottom ? '↑' : '↓'}
    </button>
  );
}
