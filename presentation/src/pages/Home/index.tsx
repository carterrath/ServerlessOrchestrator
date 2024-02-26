import { TeamSection } from '../../components/TeamSection';
import { FeaturesSection } from '../../components/FeaturesSection';
import { AboutSection } from '../../components/AboutSection';
import { ScrollAssistant } from "../../components/ScrollAssistant";
import { HeaderSection } from "../../components/HeaderSection";

export function Home() {

  return (
    <div className="flex flex-col">

      <HeaderSection />
      <AboutSection />
      <FeaturesSection />
      <TeamSection />
      
      <ScrollAssistant />

    </div>
  )
}

