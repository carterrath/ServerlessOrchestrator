export function OtherMaterialsSection(){
    return(
        <div className="bg-white text-black text-center p-10">
            <div className="bg-gradient-to-r from-selectiveYellow to-dutchWhite rounded-lg shadow-lg p-12 w-full max-w-6xl mx-auto">
                <div className="flex justify-between items-center">
                    <div className="text-left">
                        <h2 className="text-2xl font-semibold mb-4">Want a more detailed look into Serverless Orchestrator?</h2>
                        <p className="mb-6">Check out our documentation below. </p>
                    </div>
                    <div className="text-right">
                        <a href="https://docs.google.com/document/d/1ntyUV9hUDh3Jr1VFuqCIltkTqBcT-A-gCIzc0nnVt8c/edit?usp=sharing" className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-black hover:bg-gray-500 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out">
                            Documentation
                        </a>
                        <a href="#documentation" className="ml-4 inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-black hover:bg-gray-500 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out">
                            Capstone Poster
                        </a>
                    </div>
                </div>
            </div>
        </div>
    );
}