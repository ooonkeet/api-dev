const Product=require("../models/products")
const getAllProducts = async(req,res)=>{
    const {company,name,featured,sort}=req.query
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
    let apiData=Product.find(queryObject)
    if (sort){
        let sortFix=sort.replace(","," ")
        apiData=apiData.sort(sortFix)
    }
    const myData=await apiData
    res.status(200).json({myData})
}

const getAllProductsTesting=async(req,res)=>{
    const myData=await Product.find(req.query).sort("name -price") //sorting method
    // console.log(" ~ file: products.js ~line 8 ~getAllProductsTesting ~req.query",req.query);
    
    res.status(200).json({myData})
}

module.exports={getAllProducts,getAllProductsTesting}