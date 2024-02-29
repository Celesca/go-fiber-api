import { useEffect, useState } from 'react';

interface Product {
  id: number;
  name: string;
  price: number;
  imageURL: string;
}

function App() {

  const [products, setProducts] = useState<Product[]>([]);
  
  useEffect(() => {
    fetchProducts();
  }, [])

  const fetchProducts = async () => {
    try {
      const response = await fetch('/api/products')
      const data = await response.json();
      setProducts(data);
    }
    catch (error) {
      console.log(error);
    }
  }

    return (

      <div>
        <h1>Testing Text</h1>
        <h1>Product List</h1>
        {products.map((product) => (
          <div key={product.id}>
            <img src={product.imageURL}/>
            <h2>{product.name}</h2>
            <p>Price : {product.price}</p>
          </div>
        ))}

      </div>
    )
  }

export default App;