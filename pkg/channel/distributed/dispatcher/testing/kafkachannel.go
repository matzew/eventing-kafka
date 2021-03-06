package testing

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"knative.dev/eventing-contrib/kafka/channel/pkg/apis/messaging/v1beta1"
	eventingduck "knative.dev/eventing/pkg/apis/duck/v1"
	"knative.dev/pkg/apis"
)

// KafkaChannelOption enables further configuration of a KafkaChannel.
type KafkaChannelOption func(*v1beta1.KafkaChannel)

// NewKafkaChannel creates an KafkaChannel with KafkaChannelOptions.
func NewKafkaChannel(name string, namespace string, options ...KafkaChannelOption) *v1beta1.KafkaChannel {
	kafkachannel := &v1beta1.KafkaChannel{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1beta1.KafkaChannelSpec{},
	}
	for _, opt := range options {
		opt(kafkachannel)
	}
	return kafkachannel
}

func WithInitKafkaChannelConditions(kafkachannel *v1beta1.KafkaChannel) {
	kafkachannel.Status.InitializeConditions()
}

func WithKafkaChannelReady(kafkachannel *v1beta1.KafkaChannel) {
	kafkachannel.Status.MarkConfigTrue()
	kafkachannel.Status.MarkTopicTrue()
	kafkachannel.Status.MarkChannelServiceTrue()
	kafkachannel.Status.MarkServiceTrue()
	kafkachannel.Status.MarkEndpointsTrue()
	kafkachannel.Status.PropagateDispatcherStatus(&appsv1.DeploymentStatus{
		Conditions: []appsv1.DeploymentCondition{
			{
				Type:   appsv1.DeploymentAvailable,
				Status: corev1.ConditionTrue,
			},
		},
	})
}

func WithKafkaChannelAddress(a string) KafkaChannelOption {
	return func(kafkachannel *v1beta1.KafkaChannel) {
		kafkachannel.Status.SetAddress(&apis.URL{
			Scheme: "http",
			Host:   a,
		})
	}
}

func WithSubscriber(uid types.UID, uri string) KafkaChannelOption {
	return func(kafkachannel *v1beta1.KafkaChannel) {
		if kafkachannel.Spec.Subscribers == nil {
			kafkachannel.Spec.Subscribers = []eventingduck.SubscriberSpec{}
		}
		kafkachannel.Spec.Subscribers = append(kafkachannel.Spec.Subscribers, eventingduck.SubscriberSpec{
			UID: uid,
			SubscriberURI: &apis.URL{
				Scheme: "http",
				Host:   uri,
			},
		})
	}
}

func WithSubscriberReady(uid types.UID) KafkaChannelOption {
	return func(kafkachannel *v1beta1.KafkaChannel) {
		if kafkachannel.Status.SubscribableStatus.Subscribers == nil {
			kafkachannel.Status.SubscribableStatus.Subscribers = []eventingduck.SubscriberStatus{}
		}
		kafkachannel.Status.SubscribableStatus.Subscribers = append(kafkachannel.Status.SubscribableStatus.Subscribers, eventingduck.SubscriberStatus{
			Ready: corev1.ConditionTrue,
			UID:   uid,
		})
	}
}
