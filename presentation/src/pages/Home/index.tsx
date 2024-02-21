export const Home = () => {
  return (
    <div id="react-app-container">
      {/* Header Section */}
      <header className="u-bg-pink u-text-white u-p-4 u-text-center">
        <h1 className="u-m-b-2">Serverless Orchestrator</h1>
        <p>Automate your workflows in the cloud seamlessly</p>
      </header>

      {/* Features Section */}
      <section className="u-p-4">
        <div className="u-flex-c-c u-flex-wrap u-m-auto u-w-max-400">
          <div className="u-p-3 u-bg-gray-light u-m-2 u-rounded">
            <h2>Feature One</h2>
            <p>Describe the feature here.</p>
          </div>
          <div className="u-p-3 u-bg-gray-light u-m-2 u-rounded">
            <h2>Feature Two</h2>
            <p>Describe the feature here.</p>
          </div>
          <div className="u-p-3 u-bg-gray-light u-m-2 u-rounded">
            <h2>Feature Three</h2>
            <p>Describe the feature here.</p>
          </div>
          <div className="u-p-3 u-bg-gray-light u-m-2 u-rounded">
            <h2>Feature Four</h2>
            <p>Describe the feature here.</p>
          </div>
        </div>
      </section>

      {/* Call to Action Section */}
      <section className="u-bg-pink u-text-white u-p-4 u-text-center u-m-t-3">
        <h2>Get Started with Serverless Orchestrator</h2>
        <button className="btn btn-success u-m-t-2">Sign Up Now</button>
      </section>

      {/* Footer Section */}
      <footer className="u-bg-pink-dark u-text-white u-p-3 u-text-center">
        <p>&copy; {new Date().getFullYear()} Serverless Orchestrator, Inc.</p>
      </footer>
    </div>
  );
};

export default Home;
