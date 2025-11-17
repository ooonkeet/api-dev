const Product=require("../models/products")
const getAllProducts = async(req,res)=>{
    const myData=await Product.find(req.query)
    res.status(200).json({myData})
}

const getAllProductsTesting=async(req,res)=>{
    const myData=await Product.find(req.query)
    console.log(" ~ file: products.js ~line 8 ~getAllProductsTesting ~req.query",req.query);
    
    res.status(200).json({myData})
}

module.exports={getAllProducts,getAllProductsTesting}