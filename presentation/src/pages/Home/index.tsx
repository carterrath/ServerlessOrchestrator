import { TeamSection } from '../../components/TeamSection';
import { FeaturesSection } from '../../components/FeaturesSection';
import { AboutSection } from '../../components/AboutSection';
import { ScrollAssistant } from "../../components/ScrollAssistant";
import { HeaderSection } from "../../components/HeaderSection";
import { OtherMaterialsSection } from '../../components/OtherMaterialsSection';

export function Home() {

  return (
    <div className="flex flex-col">
      
      <div id="header-section" className="border-b border-slate-200">
        <HeaderSection />
      </div>
      <div id="features-section" className="border-b border-slate-200">
        <FeaturesSection />
      </div>
      <div id="about-section" >
      <AboutSection />
      </div>
      <div id="other-materials-section" >
      <OtherMaterialsSection />
      </div>
      <div id="team-section">
      <TeamSection />
      </div>
    
      
      <ScrollAssistant />

    </div>
  );
}