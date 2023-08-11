import mongoose from 'mongoose'

interface telemetryAttrs {
    time: Date
    data: number
}

interface TelemetryDocument extends mongoose.Document {
    time: Date
    data: number
}

interface TelemetryModel extends mongoose.Model<TelemetryDocument> {
    build(attrs: telemetryAttrs): TelemetryDocument
}

const telemetrySchema = new mongoose.Schema({
    time: {
        type: Date,
        required: true
    },
    data: {
        type: Number,
        required: true
    }
})

telemetrySchema.statics.build = (attrs: telemetryAttrs) => {
    return new Telemetry(attrs)
}

const Telemetry = mongoose.model<TelemetryDocument, TelemetryModel>('Telemetry', telemetrySchema)

export { Telemetry }