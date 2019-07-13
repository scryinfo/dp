const   path = require('path');

module.exports = {
    //...
    devServer: {
        contentBase: path.join(__dirname, 'public'),
        compress: false,
        port: 9000,
        headers: {
            'Access-Control-Allow-Origin': '*',
            'Access-Control-Allow-Methods':'*',
            'Access-Control-Allow-Headers':'content-type,x-grpc-web,x-user-agent'
        }
    },
    module: {
        rules:[
            {
                test: /\.html$/,
                use: [
                    // apply multiple loaders and options
                    "htmllint-loader",
                    {
                        loader: "html-loader"
                    }
                ]
            }
        ]
    }
};