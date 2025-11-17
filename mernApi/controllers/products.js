const Product=require("../models/products")
const getAllProducts = async(req,res)=>{
    const {company,name,featured}=req.query
    const queryObject={}
    if(company){
        queryObject.company=company
    }
    if(featured)
    {
        queryObject.featured=featured
    }
    if(name){
        queryObject.name={$regex: name, $options: "i"}
    }
    const myData=await Product.find(queryObject)
    res.status(200).json({myData})
}

const getAllProductsTesting=async(req,res)=>{
    const {company}=req.query
    const queryObject={}
    if(company){
        queryObject.company=company
    }
    if(featured)
    {
        queryObject.featured=featured
    }
    if(name){
        queryObject.name={$regex: name, $options: "i"}
    }
    const myData=await Product.find(queryObject)
    // console.log(" ~ file: products.js ~line 8 ~getAllProductsTesting ~req.query",req.query);
    
    res.status(200).json({myData})
}

module.exports={getAllProducts,getAllProductsTesting}