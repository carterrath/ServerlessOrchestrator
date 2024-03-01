import { TeamSection } from '../../components/TeamSection';
import { FeaturesSection } from '../../components/FeaturesSection';
import { AboutSection } from '../../components/AboutSection';
import { ScrollAssistant } from "../../components/ScrollAssistant";
import { HeaderSection } from "../../components/HeaderSection";
import { OtherMaterialsSection } from '../../components/OtherMaterialsSection';

export function Home() {

  return (
    <div className="flex flex-col">
      
      <div className="border-b border-slate-200">
        <HeaderSection />
      </div>
      <div className="border-b border-slate-200">
        <FeaturesSection />
      </div>
      
      <AboutSection />
      
      <OtherMaterialsSection />
      <TeamSection />
    
      
      <ScrollAssistant />

    </div>
  );
}