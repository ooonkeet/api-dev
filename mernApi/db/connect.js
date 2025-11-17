const mongoose = require("mongoose")
// const { options } = require("../routes/products")
uri = "mongodb+srv://ankitspare981_db_user:ClYG5Fv1VZp3NUjz@cluster0.y5lvab5.mongodb.net/?appName=Cluster0"
const connectDB=()=>{
    console.log("connect db");
    
    return mongoose.connect(uri,{
        useNewUrlParser: true,
        useUnifiedTopology: true,
    })
}

module.exports=connectDB;