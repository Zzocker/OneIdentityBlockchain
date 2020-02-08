const express = require('express')
const network = require('./contract')


const routes = express.Router()

routes.put('/admin/responRequest',async (req,res)=>{
    try {
        rbody = req.body
        const contract = await network.contract()
        resp = rbody.response
        if (resp==="-1"){
        const response = await contract.submitTransaction("responRequest",rbody.req_id,"-1")
        res.status(200).json({
            result : JSON.parse(response)
        })
    }
        if (resp==="1"){
            const response = await contract.submitTransaction("responRequest",rbody.req_id,"1",rbody.given_date)
            res.status(200).json({
                result : JSON.parse(response)
            })
        }
        
    } catch (error) {
        res.status(500).json({
            msg: error.message
        }) 
    }
})
routes.put('/admin/verify',async (req,res)=>{
    try {
        const contract = await network.contract()
        const response = await contract.submitTransaction("verifyPersonal",req.body.req_id)
        res.status(200).json({   
            books: JSON.parse(response)
        })
    } catch (error) {
        res.status(500).json({
            msg: error.message
        })
    }
})
module.exports = routes
 