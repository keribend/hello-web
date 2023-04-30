export const manifest = {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {"start":{"file":"_app/immutable/entry/start.fd0c6eb9.js","imports":["_app/immutable/entry/start.fd0c6eb9.js","_app/immutable/chunks/index.ab0e7057.js","_app/immutable/chunks/singletons.dc45313c.js"],"stylesheets":[],"fonts":[]},"app":{"file":"_app/immutable/entry/app.7c970d91.js","imports":["_app/immutable/entry/app.7c970d91.js","_app/immutable/chunks/index.ab0e7057.js"],"stylesheets":[],"fonts":[]}},
		nodes: [
			() => import('./nodes/0.js'),
			() => import('./nodes/1.js'),
			() => import('./nodes/2.js')
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0], errors: [1], leaf: 2 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		}
	}
};
